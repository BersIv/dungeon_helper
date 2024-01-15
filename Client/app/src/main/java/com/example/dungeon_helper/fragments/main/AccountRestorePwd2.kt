package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.R
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.databinding.FragmentAccountRestorePwd2Binding

class AccountRestorePwd2 : Fragment() {

    companion object {
        fun newInstance() = AccountRestorePwd2()
    }

    private lateinit var viewModel: AccountRestorePwd2ViewModel
    private var _binding: FragmentAccountRestorePwd2Binding? = null
    private val binding get() = _binding!!
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
        val accountRestorePwd2ViewModel = ViewModelProvider(this)[AccountRestorePwd2ViewModel::class.java]
        _binding = FragmentAccountRestorePwd2Binding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textRestorePwd2
        accountRestorePwd2ViewModel.text.observe(viewLifecycleOwner){
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
        val backBtn2 = binding.backBtn2
        val savePwdBtn = binding.savePwdBtn

        backBtn2.setOnClickListener{
            (activity as MainActivity).navController.navigate(R.id.action_accountRestorePwd2_to_navigation_account)
        }
        savePwdBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_accountRestorePwd2_to_navigation_account)
        }
    }

}