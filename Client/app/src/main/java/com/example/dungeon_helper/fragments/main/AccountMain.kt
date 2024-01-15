package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import com.example.dungeon_helper.MainActivity
import android.widget.TextView
import com.example.dungeon_helper.AuthActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentAccountMainBinding
import android.content.Intent

class AccountMain : Fragment() {

    companion object {
        fun newInstance() = AccountMain()
    }
    private lateinit var viewModel: AccountMainViewModel


    private var _binding: FragmentAccountMainBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val accountMainViewModel = ViewModelProvider(this)[AccountMainViewModel::class.java]

        _binding = FragmentAccountMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textAccount
        accountMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }

        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val changePwdBtn = binding.changePwdBtn
        val exAccBtn = binding.exAccBtn
        val editBtn = binding.editBtn
        changePwdBtn.setOnClickListener {
          (activity as MainActivity).navController.navigate(R.id.action_navigation_account_to_accountRestorePwd)
         }
        editBtn.setOnClickListener{
            (activity as MainActivity).navController.navigate(R.id.action_navigation_account_to_accountEdit)
        }
        exAccBtn.setOnClickListener {
            val intent = Intent(requireActivity(), AuthActivity::class.java)
            startActivity(intent)
        }
    }
}