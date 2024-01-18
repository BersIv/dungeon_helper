package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentCharacterViewingBinding

class CharacterViewing : Fragment() {

    companion object {
        fun newInstance() = CharacterViewing()
    }

    private lateinit var viewModel: CharacterViewingViewModel
    private var _binding:FragmentCharacterViewingBinding? =  null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val characterViewingViewModel = ViewModelProvider(this)[CharacterViewingViewModel::class.java]
        _binding = FragmentCharacterViewingBinding.inflate(inflater, container, false)
        val root: View = binding.root
        val textView: TextView = binding.textCharacterViewing
        characterViewingViewModel.text.observe(viewLifecycleOwner){
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
        val backBtn = binding.backBtn
        val stopViewBtn = binding.stopViewBtn
        backBtn.setOnClickListener{
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
            (activity as MainActivity).navController.navigate(R.id.action_characterViewing_to_navigation_character)

        }
        stopViewBtn.setOnClickListener {
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
            (activity as MainActivity).navController.navigate(R.id.action_characterViewing_to_navigation_character)
        }
    }

}