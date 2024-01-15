package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.databinding.FragmentAccountRestorePwdBinding
import com.example.dungeon_helper.R


class AccountRestorePwd : Fragment() {

    companion object {
        fun newInstance() = AccountRestorePwd()
    }

    private lateinit var viewModel: AccountRestorePwdViewModel
    private var _binding: FragmentAccountRestorePwdBinding? =  null
    private val binding get() = _binding !!
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val activity = requireActivity() as MainActivity
        val navView = activity.getNavView()
        val menu = navView.menu
        val menuItem1 = menu.findItem(R.id.navigation_info)
        val menuItem2 = menu.findItem(R.id.navigation_character)
        val menuItem3 = menu.findItem(R.id.navigation_lobby)
        val menuItem4 = menu.findItem(R.id.navigation_account)
        menuItem1.isVisible = false
        menuItem2.isVisible = false
        menuItem3.isVisible = false
        menuItem4.isVisible = false
        val accountRestorePwdViewModel = ViewModelProvider(this)[AccountRestorePwdViewModel::class.java]
        _binding = FragmentAccountRestorePwdBinding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textRestorePwd
        accountRestorePwdViewModel.text.observe(viewLifecycleOwner){
            textView.text = it
        }
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
        val activity = requireActivity() as MainActivity
        val navView = activity.getNavView()
        val menu = navView.menu
        val menuItem1 = menu.findItem(R.id.navigation_info)
        val menuItem2 = menu.findItem(R.id.navigation_character)
        val menuItem3 = menu.findItem(R.id.navigation_lobby)
        val menuItem4 = menu.findItem(R.id.navigation_account)
        menuItem1.isVisible = true
        menuItem2.isVisible = true
        menuItem3.isVisible = true
        menuItem4.isVisible = true
    }

    override fun onStart() {
        super.onStart()
        val backBtn = binding.backBtn
        val restoreBtn = binding.restoreBtn

        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_accountRestorePwd_to_navigation_account)
        }
        restoreBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_accountRestorePwd_to_accountRestorePwd2)
        }
    }


}